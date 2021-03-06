# This file was generated by GoReleaser. DO NOT EDIT.
class Om < Formula
  desc ""
  homepage ""
  version "6.1.1"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/pivotal-cf/om/releases/download/6.1.1/om-darwin-6.1.1.tar.gz"
    sha256 "18a7b99f6aa2cbb62ba9a40b76329635ce619f14fb7e3ab1b6233fdc06bb500b"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/pivotal-cf/om/releases/download/6.1.1/om-linux-6.1.1.tar.gz"
      sha256 "a624e7e5b4f61db67ed1a3e01fcc8de1b3f5588b76a6ce7909dd59bbccf8c326"
    end
  end

  def install
    bin.install "om"
  end

  test do
    system "#{bin}/om --version"
  end
end
