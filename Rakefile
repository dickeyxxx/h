TARGETS = [
  {os: 'darwin', arch: 'amd64'},
  {os: 'linux', arch: 'amd64'},
  {os: 'linux', arch: '386'},
  {os: 'windows', arch: 'amd64'},
  {os: 'windows', arch: '386'}
]

VERSION = File.open('VERSION').read.chomp

task :build do
  puts "Building #{VERSION}"
  TARGETS.each do |target|
    filename = "hk_#{target[:os]}_#{target[:arch]}"
    puts "Building #{filename}"
    path = "dist/releases/#{VERSION}/#{filename}"
    system("GOOS=#{target[:os]} GOARCH=#{target[:arch]} go build -o #{path}")
    system("gzip -f #{path}")
  end
end

task :deploy => :build do
  require 'aws-sdk'
  puts "Deploying #{VERSION}"
  s3 = AWS::S3.new
  bucket = s3.buckets['dickeyxxx_dev']
  TARGETS.each do |target|
    filename = "hk_#{target[:os]}_#{target[:arch]}.gz"
    puts "Uploading #{filename}"
    upload_file(bucket,  "dist/releases/#{VERSION}/#{filename}",  "releases/#{VERSION}/#{filename}")
    upload_file(bucket,  "dist/releases/#{VERSION}/#{filename}",  "releases/#{filename}")
  end
  puts "setting VERSION to #{VERSION}"
  upload_file(bucket, 'VERSION', 'releases/VERSION')
end

def upload_file(bucket, local, remote)
  obj = bucket.objects[remote]
  obj.write(Pathname.new(local))
  obj.acl = :public_read
end
