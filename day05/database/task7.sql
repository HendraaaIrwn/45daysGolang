SELECT
  p.product_id,
  p.product_name,
  COUNT(DISTINCT od.order_id) AS total_order,
  SUM(od.quantity) AS total_quantity
FROM oe.order_details od
JOIN oe.products p
  ON p.product_id = od.product_id
GROUP BY p.product_id, p.product_name
ORDER BY total_quantity DESC;
