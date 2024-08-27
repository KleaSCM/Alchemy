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
