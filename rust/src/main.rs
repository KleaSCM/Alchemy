use actix_web::{web, App, HttpResponse, HttpServer, Responder};

async fn organize() -> impl Responder {
    HttpResponse::Ok().json("Files organized")
}

async fn extract_metadata() -> impl Responder {
    HttpResponse::Ok().json("Metadata extracted")
}

async fn deduplicate() -> impl Responder {
    HttpResponse::Ok().json("Duplicates removed")
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
