#![feature(once_cell_try)]

use neon::prelude::*;

use std::sync::OnceLock;
use tokio::runtime::Runtime;

struct Arrays {
    arrays: Vec<Vec<String>>,
}

fn runtime<'a, C: Context<'a>>(cx: &mut C) -> NeonResult<&'static Runtime> {
    static RUNTIME: OnceLock<Runtime> = OnceLock::new();

    RUNTIME.get_or_try_init(|| Runtime::new().or_else(|err| cx.throw_error(err.to_string())))
}

fn arrays() -> &'static Arrays {
    static ARRAYS: OnceLock<Arrays> = OnceLock::new();

    ARRAYS.get_or_init(|| {
        let mut arrays: Vec<Vec<String>> = Vec::with_capacity(20);

        for i in 0..20 {
            if i % 2 == 0 {
                arrays.push(vec![
                    String::from("qw1"),
                    String::from("qw2"),
                    String::from("qw3"),
                    String::from("qw4"),
                    String::from("qw5"),
                ]);
            } else {
                arrays.push(vec![
                    String::from("qw5"),
                    String::from("qw4"),
                    String::from("qw3"),
                    String::from("qw2"),
                    String::from("qw1"),
                ]);
            }
        }

        Arrays { arrays }
    })
}

fn exists(mut cx: FunctionContext) -> JsResult<JsPromise> {
    let rt = runtime(&mut cx)?;
    let arrays = arrays();
    let channel = cx.channel();

    let (deferred, promise) = cx.promise();

    rt.spawn(async move {
        let mut i: u32 = 0;

        for b in 0..50000 {
            let array = &arrays.arrays[b % arrays.arrays.len()];

            for val in array {
                if val == "qw2" {
                    i += 1;

                    break;
                }
            }
        }

        tokio::time::sleep(tokio::time::Duration::from_millis(10)).await;

        for b in 0..50000 {
            let array = &arrays.arrays[b % arrays.arrays.len()];

            for val in array {
                if val == "qw5" {
                    i += 1;

                    break;
                }
            }
        }

        tokio::time::sleep(tokio::time::Duration::from_millis(10)).await;

        for b in 0..50000 {
            let array = &arrays.arrays[b % arrays.arrays.len()];

            for val in array {
                if val == "qw8" {
                    i += 1;

                    break;
                }
            }
        }

        deferred.settle_with(&channel, move |mut cx| {
            Ok(cx.number(i))
        });
    });

    Ok(promise)
}

#[neon::main]
fn main(mut cx: ModuleContext) -> NeonResult<()> {
    cx.export_function("exists", exists)?;
    Ok(())
}
