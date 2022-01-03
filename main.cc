#include <iostream>
#include <filesystem>
#include <vector>
namespace fs = std::filesystem;

class Language
{
public:
  std::string filename;
  std::vector<std::string> extensions;
  Language(std::string filename)
  {
    this->filename = filename;
    this->extensions = {"x",
                        "y",
                        "z"};
  }
  void getX() { std::cout << filename; }
};

int main31()
{
  for (const auto &entry : fs::directory_iterator("./cauldron"))
  {
    std::string currentPath = entry.path();
    std::string filename = currentPath.substr(currentPath.find_last_of("/\\") + 1);
    std::cout << filename << std::endl;
  }
  return 0;
}

int main()
{
}
