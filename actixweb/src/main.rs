use std::cmp;

use actix::prelude::*;
use actix_web::{error, get, http, web, App, HttpRequest, HttpResponse, HttpServer, Responder};
use askama::Template;
use serde::Serialize;

mod db_diesel;
mod models;
mod schema;

#[derive(Serialize)]
pub struct Message {
    pub message: &'static str,
}

fn get_query_param(query: &str) -> u16 {
    let q = if let Some(pos) = query.find('n') {
        query.split_at(pos + 2).1.parse::<u16>().ok().unwrap_or(1)
    } else {
        1
    };
    cmp::min(500, cmp::max(1, q))
}

#[get("/json")]
async fn json() -> impl Responder {
    let message = Message {
        message: "Hello, World!",
    };

    let body = serde_json::to_string(&message).unwrap();

    HttpResponse::Ok()
        .insert_header((http::header::SERVER, "Actix Web"))
        .insert_header((
            http::header::CONTENT_TYPE,
            http::header::ContentType::json(),
        ))
        .body(body)
}

#[get("/plaintext")]
async fn plaintext() -> impl Responder {
    HttpResponse::Ok()
        .insert_header((http::header::SERVER, "Actix Web"))
        .insert_header((
            http::header::CONTENT_TYPE,
            http::header::ContentType::plaintext(),
        ))
        .body("Hello, World!")
}

#[get("/db")]
async fn world_row(
    db: web::Data<Addr<db_diesel::DbExecutor>>,
) -> Result<HttpResponse, error::Error> {
    let res = db
        .send(db_diesel::RandomWorld)
        .await
        .map_err(error::ErrorInternalServerError)?;

    match res {
        Ok(row) => {
            let body = serde_json::to_string(&row).unwrap();
            Ok(HttpResponse::Ok()
                .insert_header((http::header::SERVER, "Actix Web"))
                .insert_header((
                    http::header::CONTENT_TYPE,
                    http::header::ContentType::json(),
                ))
                .body(body))
        }
        Err(_) => Ok(HttpResponse::InternalServerError().into()),
    }
}

#[get("/queries")]
async fn queries(
    req: HttpRequest,
    db: web::Data<Addr<db_diesel::DbExecutor>>,
) -> Result<HttpResponse, error::Error> {
    // get queries parameter
    let q = get_query_param(req.query_string());

    // run SQL queries
    let res = db
        .send(db_diesel::RandomWorlds(q))
        .await
        .map_err(error::ErrorInternalServerError)?;
    if let Ok(worlds) = res {
        let body = serde_json::to_string(&worlds).unwrap();
        Ok(HttpResponse::Ok()
            .insert_header((http::header::SERVER, "Actix Web"))
            .insert_header((
                http::header::CONTENT_TYPE,
                http::header::ContentType::json(),
            ))
            .body(body))
    } else {
        Ok(HttpResponse::InternalServerError().into())
    }
}

#[get("/updates")]
async fn updates(
    req: HttpRequest,
    db: web::Data<Addr<db_diesel::DbExecutor>>,
) -> Result<HttpResponse, error::Error> {
    // get queries parameter
    let q = get_query_param(req.query_string());

    // update worlds
    let res = db
        .send(db_diesel::UpdateWorld(q))
        .await
        .map_err(error::ErrorInternalServerError)?;

    if let Ok(worlds) = res {
        let body = serde_json::to_string(&worlds).unwrap();
        Ok(HttpResponse::Ok()
            .insert_header((http::header::SERVER, "Actix Web"))
            .insert_header((
                http::header::CONTENT_TYPE,
                http::header::ContentType::json(),
            ))
            .body(body))
    } else {
        Ok(HttpResponse::InternalServerError().into())
    }
}

#[derive(Template)]
#[template(path = "fortune.html")]
struct FortuneTemplate<'a> {
    items: &'a Vec<models::Fortune>,
}

#[get("/fortunes")]
async fn fortune(db: web::Data<Addr<db_diesel::DbExecutor>>) -> Result<HttpResponse, error::Error> {
    let res = db
        .send(db_diesel::TellFortune)
        .await
        .map_err(error::ErrorInternalServerError)?;
    match res {
        Ok(rows) => {
            let tmpl = FortuneTemplate { items: &rows };
            let res = tmpl.render().unwrap();

            Ok(HttpResponse::Ok()
                .insert_header((http::header::SERVER, "Actix Web"))
                .insert_header((
                    http::header::CONTENT_TYPE,
                    http::header::ContentType::html(),
                ))
                .body(res))
        }
        Err(_) => Ok(HttpResponse::InternalServerError().into()),
    }
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let db_url = "mysql://root@localhost/hello_world";

    // start DB executor actors
    let addr = SyncArbiter::start(num_cpus::get() * 3, move || {
        db_diesel::DbExecutor::new(db_url)
    });

    HttpServer::new(move || {
        App::new()
            .app_data(web::Data::new(addr.clone()))
            .service(json)
            .service(plaintext)
            .service(world_row)
            .service(queries)
            .service(updates)
            .service(fortune)
    })
    .bind("0.0.0.0:8080")?
    .run()
    .await
}
