const formatMoney = (price = 0) => {
  return price.toLocaleString('it-IT', {
    style: 'currency',
    currency: 'VND'
  });
};

export { formatMoney };
