WITH category_qty AS (
  SELECT
    c.category_id,
    c.category_name,
    SUM(od.quantity) AS total_qty_ordered
  FROM oe.order_details od
  JOIN oe.products p
    ON p.product_id = od.product_id
  JOIN oe.categories c
    ON c.category_id = p.category_id
  GROUP BY c.category_id, c.category_name
),
extremes AS (
  SELECT
    MIN(total_qty_ordered) AS min_qty,
    MAX(total_qty_ordered) AS max_qty
  FROM category_qty
)
SELECT
  cq.category_id,
  cq.category_name,
  cq.total_qty_ordered,
  CASE
    WHEN cq.total_qty_ordered = e.min_qty THEN 'MIN'
    WHEN cq.total_qty_ordered = e.max_qty THEN 'MAX'
  END AS label
FROM category_qty cq
CROSS JOIN extremes e
WHERE cq.total_qty_ordered IN (e.min_qty, e.max_qty)
ORDER BY label DESC, cq.total_qty_ordered;
