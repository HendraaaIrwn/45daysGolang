SELECT
  p.product_id,
  p.product_name,
  s.supplier_id,
  s.company_name,
  p.unit_price,
  p.units_in_stock,
  p.units_on_order,
  p.reorder_level,
  (p.reorder_level - (COALESCE(p.units_in_stock,0) + COALESCE(p.units_on_order,0))) AS qty_to_reorder
FROM oe.products p
JOIN oe.suppliers s
  ON s.supplier_id = p.supplier_id
WHERE (COALESCE(p.units_in_stock,0) + COALESCE(p.units_on_order,0)) <= COALESCE(p.reorder_level,0)
ORDER BY qty_to_reorder DESC, p.product_name;
