## DOJO - Dependency Injection - The Go way

#### Files Content
```
dependency-injection
    examples                    (trivial examples)
        no_di_pure_functions    (using functions with no dependency injection)
        third_party             (package to represent some 3rd party library)
        with_di                 (using dependency injection)
    
    practice                    (beer forecast project to exercise the dojo's content)
```

#### Practice Time
In order to improve the strong coupling and assign the right responsibility to each app layer, refactor all the
components between the handler controller and the weather forecast client using dependency injection pattern.

The aim of the exercise is to reach the point where you can replace all the package coupling by using `interface`s.
The second goal, would be to unit test the method `Estimate` from the `estimator.go` file, by mocking its dependencies.
