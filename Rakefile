require 'digest'

TARGETS = [
  {os: 'darwin', arch: 'amd64'},
  {os: 'linux', arch: 'amd64'},
  {os: 'linux', arch: '386'},
  {os: 'windows', arch: 'amd64'},
  {os: 'windows', arch: '386'}
]

VERSION = File.open('VERSION').read.chomp
COMMIT = `git rev-parse HEAD`

task :build do
  FileUtils.mkdir_p 'dist'
  puts "Building #{VERSION}"
  TARGETS.each do |target|
    path = "dist/#{VERSION}/hk_#{target[:os]}_#{target[:arch]}"
    puts "Building #{path}"
    build(target[:os], target[:arch], path)
    gzip(path)
    write_digest("#{path}.gz")
  end
end

namespace :deploy do
  task :release => :build do
    require 'aws-sdk'
    puts "Deploying #{VERSION}"
    bucket = get_s3_bucket
    TARGETS.each do |target|
      filename = "hk_#{target[:os]}_#{target[:arch]}.gz"
      puts "Uploading #{filename}"
      upload_file(bucket, "dist/#{VERSION}/#{filename}", "hk/#{VERSION}/#{filename}")
      upload_file(bucket, "dist/#{VERSION}/#{filename}.sha1", "hk/#{VERSION}/#{filename}.sha1")
      upload_file(bucket, "dist/#{VERSION}/#{filename}", "hk/#{filename}")
      upload_file(bucket, "dist/#{VERSION}/#{filename}.sha1", "hk/#{filename}.sha1")
    end
    puts "setting VERSION to #{VERSION}"
    upload_file(bucket, 'VERSION', 'hk/VERSION')
  end

  task :dev => :build do
    require 'aws-sdk'
    puts "Deploying dev"
    bucket = get_s3_bucket
    TARGETS.each do |target|
      filename = "hk_#{target[:os]}_#{target[:arch]}.gz"
      puts "Uploading #{filename}"
      upload_file(bucket, "dist/#{VERSION}/#{filename}", "hk/dev/#{filename}")
      upload_file(bucket, "dist/#{VERSION}/#{filename}.sha1", "hk/dev/#{filename}.sha1")
    end
    upload_string(bucket, COMMIT, 'hk/dev/COMMIT')
  end
end

def build(os, arch, path)
  system("GOOS=#{os} GOARCH=#{arch} go build -o #{path}")
end

def gzip(path)
  system("gzip -f #{path}")
end

def write_digest(path)
  digest = Digest::SHA1.file(path).hexdigest
  File.open(path + '.sha1', 'w') { |f| f.write(digest) }
end

def get_s3_bucket
  s3 = AWS::S3.new
  s3.buckets['dickeyxxx_dev']
end

def upload_file(bucket, local, remote)
  obj = bucket.objects[remote]
  obj.write(Pathname.new(local))
  obj.acl = :public_read
end

def upload_string(bucket, s, remote)
  obj = bucket.objects[remote]
  obj.write(s)
  obj.acl = :public_read
end
