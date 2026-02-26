import org.apache.spark.ml.{Pipeline, PipelineModel}
import org.apache.spark.ml.classification.LogisticRegression
import org.apache.spark.ml.feature.{HashingTF, Tokenizer}
import org.apache.spark.sql.{Row, SparkSession}
import org.apache.spark.ml.linalg.Vector

object PipelineExample {
  def main(args: Array[String]): Unit = {
    val spark = SparkSession
      .builder()
      .appName("PipelineExample")
      .getOrCreate()
    val training = spark.createDataFrame(Seq(
      (0L, "a b c d e spark", 1.0),
      (1L, "b d", 0.0),
      (2L, "spark f g h", 1.0),
      (3L, "hadoop mapreduce", 0.0)
    )).toDF("id", "text", "label")
    //Configure an ML pipeline, which consists of three stages: tokenizer, hashingTF and lr
    val tokenizer = new Tokenizer()
      .setInputCol("text")
      .setOutputCol("words")
    val hashingTF = new HashingTF()
      .setNumFeatures(1000)
      .setInputCol(tokenizer.getOutputCol)
      .setOutputCol("features")
    val lr = new LogisticRegression()
      .setMaxIter(10)
      .setRegParam(0.001)
    val pipeline = new Pipeline()
      .setStages(Array(tokenizer, hashingTF, lr))
    //Fit the pipeline to training documents
    val model = pipeline.fit(training)
    //Now we can optionally save the fitted pipeline to disk
    model.write.overwrite().save("/tmp/spark-logistic-regression-model")
    //We can also save this unfit pipeline to disk
    pipeline.write.overwrite().save("/tmp/unfit-lr-model")
    //And load it back in during production
    //The code fails if the directory /tmp/spark-logistic-regrssion-model has not been created yet
    val sameModel = PipelineModel.load("/tmp/spark-logistic-regression-model")
    //Prepare test documents, which are unlabeled (id, text) tuples
    val test = spark.createDataFrame(Seq(
      (4L, "spark i j k"),
      (5L, "l m n"),
      (6L, "spark hadoop spark"),
      (7L, "apache hadoop")
    )).toDF("id", "text")
    //Make predictions on test documents
    model.transform(test)
      .select("id", "text", "probability", "prediction")
      .collect()
      .foreach { case Row(id: Long, text: String, prob: Vector, prediction: Double) =>
        println(s"($id, $text)-->prob=$prob, prediction=$prediction")
      }
    spark.stop()
  }
}
/////////////////////////////////////////////////////////////////////////////////
// Output
/////////////////////////////////////////////////////////////////////////////////
/* (4, spark i j k)-->prob=[0.6292098489668487,0.37079015103315127], prediction=0.0
(5, l m n)-->prob=[0.984770006762304,0.015229993237696027], prediction=0.0
(6, spark hadoop spark)-->prob=[0.13412348342566105,0.865876516574339], prediction=1.0
(7, apache hadoop)-->prob=[0.9955732114398529,0.00442678856014711], prediction=0.0 */
