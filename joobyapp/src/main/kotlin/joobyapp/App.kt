package joobyapp

import com.fasterxml.jackson.databind.ObjectMapper
import io.jooby.Context
import io.jooby.ExecutionMode.EVENT_LOOP
import io.jooby.MediaType.JSON
import io.jooby.exception.BadRequestException
import io.jooby.hikari.HikariModule
import io.jooby.jackson.JacksonModule
import io.jooby.kt.require
import io.jooby.kt.runApp
import io.jooby.rocker.RockerModule
import java.util.*
import java.util.concurrent.ThreadLocalRandom
import javax.sql.DataSource

private const val SELECT_WORLD = "select * from world where id=?"
private const val UPDATE_WORLD = "update world SET randomNumber=? where id=?"

private const val DB_ROWS = 10000

private const val MESSAGE = "Hello, World!"

private val MESSAGE_BYTES = MESSAGE.toByteArray(Charsets.UTF_8)

data class Message(val message: String = MESSAGE)

class World(val id: Int, var randomNumber: Int)

fun main(args: Array<String>) {
  runApp(args, EVENT_LOOP) {

    /** Template engine: */
    install(RockerModule().reuseBuffer(true))

    /** JSON: */
    install(JacksonModule())
    val mapper = require(ObjectMapper::class)

    /** Database: */
    install(HikariModule())
    val ds = require(DataSource::class)

    get("/plaintext") { ctx.send(MESSAGE_BYTES) }

    get("/json") {
      ctx.setResponseType(JSON)
      ctx.send(mapper.writeValueAsBytes(Message()))
    }

    /** Go blocking : */
    dispatch {

      /** Single query: */
      get("/db") {
        val rnd = ThreadLocalRandom.current()
        val result =
            ds.connection.use { conn ->
              conn.prepareStatement(SELECT_WORLD).use { statement ->
                statement.setInt(1, nextRandom(rnd))
                statement.executeQuery().use { rs ->
                  rs.next()
                  World(rs.getInt(1), rs.getInt(2))
                }
              }
            }
        ctx.setResponseType(JSON)
        ctx.send(mapper.writeValueAsBytes(result))
      }

      /** Multiple query: */
      get("/queries") {
        val rnd = ThreadLocalRandom.current()
        val queries = ctx.queries()
        val result = ArrayList<World>(queries)
        ds.connection.use { conn ->
          conn.prepareStatement(SELECT_WORLD).use { statement ->
            for (i in 1..queries) {
              statement.setInt(1, nextRandom(rnd))
              statement.executeQuery().use { rs ->
                while (rs.next()) {
                  result += World(rs.getInt(1), rs.getInt(2))
                }
              }
            }
          }
        }
        ctx.setResponseType(JSON)
        ctx.send(mapper.writeValueAsBytes(result))
      }

      /** Updates: */
      get("/updates") {
        val queries = ctx.queries()
        val result = ArrayList<World>(queries)
        val rnd = ThreadLocalRandom.current()

        ds.connection.use { connection ->
          connection.prepareStatement(SELECT_WORLD).use { statement ->
            for (i in 1..queries) {
              statement.setInt(1, nextRandom(rnd))
              statement.executeQuery().use { rs ->
                rs.next()
                result += World(rs.getInt("id"), rs.getInt("randomNumber"))
              }
            }
          }

          for (world in result) {
            world.randomNumber = nextRandom(rnd)
            connection.prepareStatement(UPDATE_WORLD).use { statement ->
              statement.setInt(1, world.id)
              statement.setInt(2, world.randomNumber)
              statement.executeUpdate()
            }
          }
        }
        ctx.setResponseType(JSON)
        ctx.send(mapper.writeValueAsBytes(result))
      }

      /** Fortunes: */
      get("/fortunes") {
        val fortunes = ArrayList<Fortune>()
        ds.connection.use { connection ->
          connection.prepareStatement("select * from fortune").use { stt ->
            stt.executeQuery().use { rs ->
              while (rs.next()) {
                fortunes += Fortune(rs.getInt("id"), rs.getString("message"))
              }
            }
          }
        }
        fortunes.add(Fortune(0, "Additional fortune added at request time."))
        fortunes.sortBy { it.message }

        /** render view: */
        views.fortunes.template(fortunes)
      }
    }
  }
}

private fun nextRandom(rnd: Random): Int {
  return rnd.nextInt(DB_ROWS) + 1
}

fun Context.queries() =
    try {
      this.query("n").intValue(1).coerceIn(1, 500)
    } catch (x: BadRequestException) {
      1
    }