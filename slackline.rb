class Slackline < Formula
  desc "Post updates/messages to slack with minimal disruption from the cmd line."
  homepage "https://github.com/davidhampgonsalves/slackline"
  url "https://github.com/davidhampgonsalves/slackline"
  version "1.0"

  depends_on "go" => :build
  depends_on "gpm" => :build

  def install
    ENV["GOPATH"] = buildpath
    system "go", "get", "github.com/davidhampgonsalves/slackline"

    # Install Go dependencies
    # system "gpm", "install"

    # Build and install termshare
    # system "go", "build", "-o", "slackline"
    bin.install "slackline"
  end

  test do
    system "#{bin}/slackline", "--version"
  end
end
