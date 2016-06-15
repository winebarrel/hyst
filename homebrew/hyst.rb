require 'formula'

class Hyst < Formula
  VERSION = '0.1.0'

  homepage 'https://github.com/winebarrel/hyst'
  url "https://github.com/winebarrel/hyst/releases/download/v#{VERSION}/hyst-v#{VERSION}-darwin-amd64.gz"
  sha256 'ba99b5ec8016eae462671b51b2c754f38d7194db8d66b1a58ee6b375000c6840'
  version VERSION
  head 'https://github.com/winebarrel/hyst.git', :branch => 'master'

  def install
    system "mv hyst-v#{VERSION}-darwin-amd64 hyst"
    bin.install 'hyst'
  end
end
