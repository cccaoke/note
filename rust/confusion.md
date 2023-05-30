## 生命周期

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

## 闭包

如下函数无法通过编译，需要将update_string修改为let mut update_string.
```rust
fn main() {
    let mut s = String::new();

    let update_string =  |str| s.push_str(str);
    update_string("hello");

    println!("{:?}",s);
}
```

但是下面的函数却可以正常通过编译

```rust
fn main() {
    let mut s = String::new();

    let update_string =  |str| s.push_str(str);

    exec(update_string);

    println!("{:?}",s);
}

fn exec<'a, F: FnMut(&'a str)>(mut f: F)  {
    f("hello")
}
```
有一种解释是说exec函数拿走了update_string的所有权，这时和update_string本身是否可变就没有关系了。
