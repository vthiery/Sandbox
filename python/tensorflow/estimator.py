import tensorflow as tf
import numpy as np

def train(estimator, input_fn, train_input_fn):
    estimator.train(input_fn=input_fn, steps=1000)
    train_metrics = estimator.evaluate(input_fn=train_input_fn)
    print("train metrics: %r"% train_metrics)

def evaluate(estimator, eval_input_fn):
    eval_metrics = estimator.evaluate(input_fn=eval_input_fn)
    print("eval metrics: %r"% eval_metrics)

# Training set
x_train = np.array([1., 2., 3., 4.])
y_train = np.array([0., -1., -2., -3.])
# Evaluation set
x_eval = np.array([2., 5., 8., 1.])
y_eval = np.array([-1.01, -4.1, -7, 0.])

input_fn = tf.estimator.inputs.numpy_input_fn(
    {"x": x_train}, y_train, batch_size=4, num_epochs=None, shuffle=True)
train_input_fn = tf.estimator.inputs.numpy_input_fn(
    {"x": x_train}, y_train, batch_size=4, num_epochs=1000, shuffle=False)

eval_input_fn = tf.estimator.inputs.numpy_input_fn(
    {"x": x_eval}, y_eval, batch_size=4, num_epochs=1000, shuffle=False)

# Linear estimator
linear_estimator = tf.estimator.LinearRegressor(
    feature_columns=[tf.feature_column.numeric_column("x", shape=[1])])

train(linear_estimator, input_fn, train_input_fn)
evaluate(linear_estimator, eval_input_fn)

###

def my_model(features, labels, mode):
    W = tf.get_variable('W', [1], dtype=tf.float64)
    b = tf.get_variable('b', [1], dtype=tf.float64)
    y = W * features['x'] + b

    loss = tf.reduce_sum(tf.square(y - labels))

    global_step = tf.train.get_global_step()
    optimizer = tf.train.GradientDescentOptimizer(0.01)
    train = tf.group(optimizer.minimize(loss), tf.assign_add(global_step, 1))

    return tf.estimator.EstimatorSpec(
        mode=mode,
        predictions=y,
        loss=loss,
        train_op=train)

custom_estimator = tf.estimator.Estimator(model_fn=my_model)

# Custom estimator
train(custom_estimator, input_fn, train_input_fn)
evaluate(custom_estimator, eval_input_fn)