from conans import ConanFile

class Sandbox(ConanFile):
    requires = "jsonformoderncpp/2.1.1@vthiery/stable"
    generators = "cmake", "gcc"

    def imports(self):
        self.copy("*.h", dst="include", src="include")

