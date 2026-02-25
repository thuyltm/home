import org.apache.spark.ml.classification.LogisticRegression
import org.apache.spark.sql.{Row, SparkSession}
import org.apache.spark.ml.linalg.Vectors
import org.apache.spark.ml.linalg.Vector
import org.apache.spark.ml.param.ParamMap

object EstimatorTransformerParamExample {
  def main(args: Array[String]): Unit = {
    val spark = SparkSession
        .builder()
        .appName("EstimatorTransformerParamExample")
        .getOrCreate()
    //Prepare training data from a list of (label, features) tuples
    val training = spark.createDataFrame(Seq(
      (1.0, Vectors.dense(0.0, 1.1, 0.1)),
      (0.0, Vectors.dense(2.0, 1.0, -1.0)),
      (0.0, Vectors.dense(2.0, 1.3, 1.0)),
      (1.0, Vectors.dense(0.0, 1.2, -0.5))
    )).toDF("label", "features")
    //Create a LogisticRegression instance. This instance is an Estimator
    val lr = new LogisticRegression()
    //Print out the parameters, documentation, and any default values
    println(s"LogisticRegression parameters:\n ${lr.explainParams()}\n")
    //We may set parameters using setter methods
    lr.setMaxIter(10)
      .setRegParam(0.01)
    //Learn a LogisticRegression model. This uses the parameters stored in lr
    val model1 = lr.fit(training)
    //Since model1 is a Model (i.e., a Transformer produced by an Estimator),
    //we can view the parameters it used during fit().
    //This prints the parameter (name: value) pairs, where names are unique IDs for this
    //LogisticRegression instance
    println(s"Model 1 was fit using parameters: ${model1.parent.extractParamMap()}")
    //We may alternatively specify parameters using a ParaMap,
    //which supports several methods for specifying parameters
    val paramMap = ParamMap(lr.maxIter->20)
      .put(lr.maxIter, 30)//Specify 1 Param. This overwrites the original maxIter.
      .put(lr.regParam->0.1, lr.threshold->0.55)//Specify multiple Params
    //One can also combine ParamMaps
    val paramMap2 = ParamMap(lr.probabilityCol->"myProbability")//Change output column name
    val paramMapCombined = paramMap ++ paramMap2
    //Now learn a new model using the paramMapCombined parameters
    //paramMapCombined overrides all parameters set earlier via lr.set* methods.
    val model2 = lr.fit(training, paramMapCombined)
    println(s"Model 2 was fit using parameters: ${model2.parent.extractParamMap()}")
    //Prepare test data
    val test = spark.createDataFrame(Seq(
      (1.0, Vectors.dense(-1.0, 1.5, 1.3)),
      (0.0, Vectors.dense(3.0, 2.0, -0.1)),
      (1.0, Vectors.dense(0.0, 2.2, -1.5))
    )).toDF("label", "features")
    //Make predictions on test data using the Transformer.transform() method.
    //LogisticRegression.transform will only use the 'features' column.
    //Note that model2.transform() outputs a 'myProbability' column instead of the usual
    //'probability' column since we renamed the lr.probabilityCol parameter previously
    model2.transform(test)
      .select("features", "label", "myProbability", "prediction")
      .collect()
      .foreach { case Row(features: Vector, label: Double, prob: Vector, prediction: Double) =>
        println(s"($features, $label)->prob=$prob, prediction=$prediction")
      }
    spark.stop()
  }
}

///////////////////////////////////////////////////////////////
// Output
//////////////////////////////////////////////////////////////
/* 
Model 1 which is LogisticRegression model was fit using parameters: {
  logreg_d830a6547cfe-aggregationDepth: 2,
  logreg_d830a6547cfe-elasticNetParam: 0.0,
  logreg_d830a6547cfe-family: auto,
  logreg_d830a6547cfe-featuresCol: features,
  logreg_d830a6547cfe-fitIntercept: true,
  logreg_d830a6547cfe-labelCol: label,
  logreg_d830a6547cfe-maxBlockSizeInMB: 0.0,
  logreg_d830a6547cfe-maxIter: 10,
  logreg_d830a6547cfe-predictionCol: prediction,
  logreg_d830a6547cfe-probabilityCol: probability,
  logreg_d830a6547cfe-rawPredictionCol: rawPrediction,
  logreg_d830a6547cfe-regParam: 0.01,
  logreg_d830a6547cfe-standardization: true,
  logreg_d830a6547cfe-threshold: 0.5,
  logreg_d830a6547cfe-tol: 1.0E-6
}
Model 2 update model 1 by tunning key hyperparameters -- maxIter, regParam, threshold -- was fit using parameters: {
  logreg_d830a6547cfe-aggregationDepth: 2,
  logreg_d830a6547cfe-elasticNetParam: 0.0,
  logreg_d830a6547cfe-family: auto,
  logreg_d830a6547cfe-featuresCol: features,
  logreg_d830a6547cfe-fitIntercept: true,
  logreg_d830a6547cfe-labelCol: label,
  logreg_d830a6547cfe-maxBlockSizeInMB: 0.0,
  logreg_d830a6547cfe-maxIter: 30,
  logreg_d830a6547cfe-predictionCol: prediction,
  logreg_d830a6547cfe-probabilityCol: myProbability,
  logreg_d830a6547cfe-rawPredictionCol: rawPrediction,
  logreg_d830a6547cfe-regParam: 0.1,
  logreg_d830a6547cfe-standardization: true,
  logreg_d830a6547cfe-threshold: 0.55,
  logreg_d830a6547cfe-tol: 1.0E-6
}
([-1.0,1.5,1.3], 1.0)->prob=[0.0570730499357254,0.9429269500642746], prediction=1.0
([3.0,2.0,-0.1], 0.0)->prob=[0.9238521956443227,0.07614780435567725], prediction=0.0
([0.0,2.2,-1.5], 1.0)->prob=[0.10972780286187782,0.8902721971381222], prediction=1.0
*/