# boil

a touch wrapper for programmers that generates boilerplate files.

to use:

`boil <OPTIONS> <ARGUMENTS>`
___
for example:

`boil Name.java`:

```java
class Name {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
    }
}
```

`boil app.py`:

```python
def main():
    print("Hello world")


if __name__ == "__main__":
    main()
```

`boil app.cc`:

```cpp
#include <iostream>

int main()
{
    printf("Hello World!");
    return 0;
}

```
