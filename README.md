# Test Driven Development
- Developed as a part of Extreme Programming
- Three steps:
  - Write a test for the next bit of functionality you want to add.
  - Write the functional code until the test passes.
  - Refactor both new and old code to make it well structured.
- Follow "Red, Green, Refactor" cycle.
  - Think &rarr; Think of a small test that will require your new code to pass.
  - Red bar &rarr; Write the test. Respect encapsulation and use the classes/interfaces that don't exist yet. This forces you to design the interfaces from the user's perspective.
  - Green bar &rarr; Write enough production code to pass the test. ~< 5lines. Hardcoding is fine. This is compare your intent with reality.
  - Refactor &rarr; Refactor and fix the code without worrying about breaking anything. No longer than 5 minutes.
- The key is tiny increments. This is to make sure every step is tested.
- Think for a test for which the code you will be writing is the fix. Even if some other logic could satisfy the test, the test should use the new code to work.
- If a feature requires two tests, write one test, implement code for just that test, then repeat.

## Reference
- https://martinfowler.com/bliki/TestDrivenDevelopment.html
- http://www.jamesshore.com/v2/books/aoad1/test_driven_development
