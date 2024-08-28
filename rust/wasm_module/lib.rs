use wasm_bindgen::prelude::*;

// testfn
#[wasm_bindgen]
pub fn greet(name: &str) -> String {
    format!("Hello, {}!", name)
}
