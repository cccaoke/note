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

### example1
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

### example2
另一个比较疑惑的例子：  
```rust
fn fn_once<F>(func: F)
where
    F: FnOnce(usize) -> bool,
{
    println!("{}", func(3));
    println!("{}", func(4));
}

fn main() {
    let x = vec![1, 2, 3];
    fn_once(|z|{z == x.len()})
}

```

上面这段代码是无法通过编译的，编译报错如下：  
```rust
   Compiling playground v0.0.1 (/playground)
error[E0382]: use of moved value: `func`
 --> src/main.rs:7:20
  |
2 | fn fn_once<F>(func: F)
  |               ---- move occurs because `func` has type `F`, which does not implement the `Copy` trait
...
6 |     println!("{}", func(3));
  |                    ------- `func` moved due to this call
7 |     println!("{}", func(4));
  |                    ^^^^ value used here after move
  |
note: this value implements `FnOnce`, which causes it to be moved when called
 --> src/main.rs:6:20
  |
6 |     println!("{}", func(3));
  |                    ^^^^
help: consider further restricting this bound
  |
4 |     F: FnOnce(usize) -> bool + Copy,
  |                              ++++++

For more information about this error, try `rustc --explain E0382`.
error: could not compile `playground` due to previous error
```

其中说了如果一个闭包实现了FnOnce，当他被调用时就失去了所有权，只有其再实现了Copy特征才能重复调用。

## 类型转换
### example1
```rust
fn main() {
    let arr :[u64; 13] = [0; 13];
    assert_eq!(std::mem::size_of_val(&arr), 8 * 13);
    let a: *const [u64] = &arr;
    let b = a as *const [u8];
    unsafe {
        assert_eq!(std::mem::size_of_val(&*b), 13)
    }
}
```
没看懂为什么b变成了13，而不是8*13，看起来b中所有元素都被修改为了u8？