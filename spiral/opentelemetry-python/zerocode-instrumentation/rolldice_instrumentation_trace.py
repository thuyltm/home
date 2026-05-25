from random import randint
from flask import Flask
from opentelemetry import trace

tracer = trace.get_tracer("diceroller.tracer")
app = Flask(__name__)

@app.route("/rolldice")
def roll_dice():
    return str(roll())

def roll():
    with tracer.start_as_current_span("roll") as rollspan:
        res = randint(1, 6)
        rollspan.set_attribute("roll.value", res)
        return res