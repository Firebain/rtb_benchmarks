use axum::{extract::State, routing::get, Json, Router};
use std::sync::Arc;

struct Arrays {
    arrays: Vec<Vec<String>>,
}

#[derive(serde::Serialize)]
struct Result {
    result: u32,
}

async fn handler(State(arrays): State<Arc<Arrays>>) -> Json<Result> {
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

    Json(Result { result: i })
}

#[tokio::main]
async fn main() {
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

    let arrays = Arc::new(Arrays { arrays });

    let app = Router::new().route("/", get(handler)).with_state(arrays);

    axum::Server::bind(&"0.0.0.0:8080".parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}
