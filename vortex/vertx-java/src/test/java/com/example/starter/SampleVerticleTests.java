package com.example.starter;

import io.vertx.core.Vertx;
import io.vertx.ext.web.client.WebClient;
import io.vertx.ext.web.codec.BodyCodec;
import io.vertx.junit5.Checkpoint;
import io.vertx.junit5.VertxExtension;
import io.vertx.junit5.VertxTestContext;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;

import java.util.concurrent.atomic.AtomicInteger;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

@DisplayName("👋 A fairly basic test example")
@ExtendWith(VertxExtension.class)
public class SampleVerticleTests {
    @Test
    @DisplayName("⏱ Count 3 timer ticks")
    void countThreeTicks(Vertx vertx, VertxTestContext testContext) {
        AtomicInteger counter = new AtomicInteger();
        vertx.setPeriodic(100, id -> {
            if (counter.incrementAndGet() == 3) {
                testContext.completeNow();
            }
        });
    }
    @Test
    @DisplayName("⏱ Count 3 timer ticks, with a checkpoint")
    void countThreeTicksWithCheckpoints(Vertx vertx, VertxTestContext testContext) {
        Checkpoint checkpoint = testContext.checkpoint(3);
        vertx.setPeriodic(100, id -> checkpoint.flag());
    }
    @Test
    @DisplayName("🚀 Deploy a HTTP service verticle and make 10 requests")
    void useSampleVerticle(Vertx vertx, VertxTestContext testContext) {
        WebClient webClient = WebClient.create(vertx);
        Checkpoint deploymentCheckpoint = testContext.checkpoint();
        Checkpoint requestCheckpoint = testContext.checkpoint(10);

        vertx.deployVerticle(new MainVerticle()).onComplete(testContext.succeeding(id -> {
            deploymentCheckpoint.flag();

            for (int i = 0; i < 10; i++) {
                webClient.get(8888, "localhost", "/")
                        .as(BodyCodec.string())
                        .send()
                        .onComplete(testContext.succeeding(resp -> {
                            testContext.verify(() -> {
                                assertEquals(200, resp.statusCode());
                                System.out.println(resp.body());
                                assertTrue(resp.body().contains("Hello from Vert.x!"));
                                requestCheckpoint.flag();
                            });
                        }));
            }
        }));
    }

}
