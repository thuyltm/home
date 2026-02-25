import org.apache.spark.sql.SparkSession

object SimpleApp {
  def main(args: Array[String]): Unit = {
    //Create a SparkSession - the entry point to Spark functionality
    val spark = SparkSession.builder
      .appName("SimpleApp")
      .master("local[*]")//Run Spark in local mode, using all available
      .getOrCreate()
    //Import implicit conversion for RDDs, DataFrames, etc.
    import spark.implicits._
    //Create a simple DataFrame and perform an action (show)
    val df = Seq("Hello", "world", "from", "Apache", "Spark").toDF("word")
    println("The DataFrame content is:")
    df.show()
    //Count the number of line containing `a` and the number containing `b` in the Spark README
    //val logFile = getClass.getResourceAsStream("/README.md")
    val filePath = "src/main/resources/README.md"
    val logData = spark.read.textFile(filePath).cache()
    val numAs = logData.filter(line=>line.contains("a")).count()
    val numBs = logData.filter(line=>line.contains("b")).count()
    println(s"Lines with a: $numAs, Lines with b: $numBs")
    //Stop the SparkSession
    spark.stop()
  }
}