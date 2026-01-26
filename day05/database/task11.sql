WITH shipper_product_qty AS (
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
  WHERE o.shipped_date IS NOT NULL
  GROUP BY
    s.shipper_id, s.company_name,
    p.product_id, p.product_name
),
ranked AS (
  SELECT
    *,
    RANK() OVER (PARTITION BY shipper_id ORDER BY total_qty_ordered DESC) AS r_max,
    RANK() OVER (PARTITION BY shipper_id ORDER BY total_qty_ordered ASC)  AS r_min
  FROM shipper_product_qty
)
SELECT
  shipper_id,
  company_name,
  product_id,
  product_name,
  total_qty_ordered
FROM ranked
WHERE r_max = 1 OR r_min = 1
ORDER BY shipper_id, total_qty_ordered;
