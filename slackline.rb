class Slackline < Formula
  desc "Post updates/messages to slack with minimal disruption from the cmd line."
  homepage "https://github.com/davidhampgonsalves/slackline"
  url "https://github.com/davidhampgonsalves/slackline/archive/1.0.tar.gz"
  version "1.0"

  def install
    bin.install "slackline-osx" => "slackline"
  end

  test do
    system "#{bin}/slackline", "--version"
  end
end
