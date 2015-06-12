class Slackline < Formula
  desc "Post updates/messages to slack with minimal disruption from the cmd line."
  homepage "https://github.com/davidhampgonsalves/slackline"
  url "https://github.com/davidhampgonsalves/slackline/releases/download/1.0/slackline"
  version "1.0"

  def install
    bin.install "slackline"
  end

  test do
    system "#{bin}/slackline", "--version"
  end
end
