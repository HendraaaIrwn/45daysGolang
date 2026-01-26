SELECT
  s.supplier_id,
  s.company_name AS supplier,
  COUNT(p.product_id) AS total_product,
  to_char(AVG(p.unit_price)::numeric, 'FM999G999G999G990D00') AS avg_unit_price
FROM oe.suppliers s
LEFT JOIN oe.products p
  ON p.supplier_id = s.supplier_id
GROUP BY s.supplier_id, s.company_name
ORDER BY total_product DESC, supplier;
