import tensorflow as tf

sess = tf.Session()

# Mere addition on the graph
a = tf.placeholder(tf.float32)
b = tf.placeholder(tf.float32)
sum = a + b

print(sess.run(sum, {a: 4.2, b: 1.3}))
print(sess.run(sum, {a: [1, 4], b: [3, 2]}))
