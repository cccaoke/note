- 生命周期

```rust
#[derive(Debug)]
struct Foo;

impl Foo {
fn mutate_and_share(&mut self) -> &Self {
&*self
}
fn share(&self) {}
}

fn main() {
let mut foo = Foo;
let loan = foo.mutate_and_share();
foo.share();
println!("{:?}", loan);
}
```
上述代码无法通过编译，因为loan的生命周期跟mutate_and_share一致，导致了一直在进行可变借用。
