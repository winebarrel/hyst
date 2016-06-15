require 'formula'

class Hyst < Formula
  VERSION = '0.1.0'

  homepage 'https://github.com/winebarrel/hyst'
  url "https://github.com/winebarrel/hyst/releases/download/v#{VERSION}/hyst-v#{VERSION}-darwin-amd64.gz"
  sha256 '...'
  version VERSION
  head 'https://github.com/winebarrel/hyst.git', :branch => 'master'

  def install
    system "mv hyst-v#{VERSION}-darwin-amd64 hyst"
    bin.install 'hyst'
  end
end
