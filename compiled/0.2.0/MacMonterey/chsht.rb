class Chsht < Formula
	homepage "https://github.com/mvrpl/Terminal-Cheat-Sheet"
	url "https://github.com/mvrpl/Terminal-Cheat-Sheet/blob/master/compiled/0.2.0/MacMonterey/chsht?raw=true"
	sha256 "85040d5c0997057d37ea7364869224cac54931b2aee7a0f41e272708cb96d4fd"
	depends_on "less"
	depends_on macos: :monterey
	depends_on arch: :x86_64
	def install
		bin.install "chsht"
	end
end
