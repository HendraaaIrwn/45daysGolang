SELECT
  s.shipper_id,
  s.company_name,
  p.product_id,
  p.product_name,
  SUM(od.quantity) AS total_qty_ordered
FROM oe.orders o
JOIN oe.shippers s
  ON s.shipper_id = o.ship_via
JOIN oe.order_details od
  ON od.order_id = o.order_id
JOIN oe.products p
  ON p.product_id = od.product_id
WHERE o.shipped_date IS NOT NULL   -- yang benar-benar sudah dikirim
GROUP BY
  s.shipper_id, s.company_name,
  p.product_id, p.product_name
ORDER BY
  s.shipper_id, total_qty_ordered DESC;
