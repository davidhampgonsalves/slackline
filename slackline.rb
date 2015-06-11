class Slackline < Formula
  desc "Post updates/messages to slack with minimal disruption from the cmd line."
  homepage "https://github.com/davidhampgonsalves/slackline"
  url "https://github.com/davidhampgonsalves/slackline/archive/1.0.tar.gz"
https://github.com/wincent/clipper/archive/0.2.tar.gz
  version "1.0"

  depends_on "go" => :build
  depends_on "gpm" => :build

  def install
    ENV["GOPATH"] = buildpath
    system "gpm", "install"
    system "go", "build", "slackline.go"

    bin.install "slackline"
  end

  test do
    system "#{bin}/slackline", "--version"
  end
end
