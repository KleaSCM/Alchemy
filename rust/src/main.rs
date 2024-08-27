use actix_web::{web, App, HttpResponse, HttpServer, Responder};
use sha2::{Sha256, Digest};
use std::collections::HashMap;
use std::fs;
use std::path::Path;

async fn organize() -> impl Responder {
    HttpResponse::Ok().json("Files organized")
}

async fn extract_metadata() -> impl Responder {
    HttpResponse::Ok().json("Metadata extracted")
}

async fn deduplicate() -> impl Responder {
    let directory = r"C:\Users\Kliea\Documents\TestTESTtestALCHEMY"; // Use raw string literal to avoid escaping issues
    let mut hash_map: HashMap<String, String> = HashMap::new();
    let mut duplicates = Vec::new();

    if let Err(e) = fs::read_dir(directory).map(|entries| {
        for entry in entries {
            let entry = entry.unwrap();
            let path = entry.path();
            if path.is_file() {
                let content = fs::read(&path).unwrap();
                let mut hasher = Sha256::new();
                hasher.update(&content);
                let checksum = format!("{:x}", hasher.finalize());

                if let Some(original_path) = hash_map.get(&checksum) {
                    duplicates.push(path.to_string_lossy().to_string());
                    duplicates.push(original_path.clone());
                } else {
                    hash_map.insert(checksum, path.to_string_lossy().to_string());
                }
            }
        }
    }) {
        return HttpResponse::InternalServerError().json(format!("Failed to read directory: {:?}", e));
    }

    HttpResponse::Ok().json(duplicates)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .route("/organize", web::post().to(organize))
            .route("/metadata", web::get().to(extract_metadata))
            .route("/deduplicate", web::post().to(deduplicate))
    })
    .bind("127.0.0.1:8081")?
    .run()
    .await
}
