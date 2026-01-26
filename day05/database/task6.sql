SELECT
  o.order_id,
  c.customer_id,
  o.order_date,
  o.required_date,
  o.shipped_date,
  (o.shipped_date - o.order_date) AS delivery_time
FROM oe.orders o
JOIN oe.customers c
  ON c.customer_id = o.customer_id
WHERE o.shipped_date IS NOT NULL
  AND (o.shipped_date - o.order_date) > 7
ORDER BY delivery_time DESC, o.order_date;
