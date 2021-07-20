SELECT
  orders.id,
  orders.status,
  orders.total_price,
  orders.created_at,
  order_items.product_id,
  order_items.quantity
FROM
  orders
  INNER JOIN order_items ON orders.id = order_items.order_id
where
  user_id = 2
  and status = "Đã thanh toán"
