use actix_web::{web, App, HttpResponse, HttpServer, Responder};

async fn organize() -> impl Responder {
    HttpResponse::Ok().json("Files organized")
}