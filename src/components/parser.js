class PaymentParser {
  /**
   * @param {string} paymentString
   */
  constructor(paymentString) {
    this.paymentString = paymentString;
  }

  parse() {
    try {
      const paymentData = JSON.parse(this.paymentString);
      return this.validatePaymentData(paymentData);
    } catch (error) {
      throw new Error('Invalid payment string');
    }
  }

  /**
   * @param {object} paymentData
   */
  validatePaymentData(paymentData) {
    if (!paymentData || typeof paymentData !== 'object') {
      throw new Error('Invalid payment data');
    }

    const requiredFields = ['amount', 'currency', 'payer', 'payee'];
    for (const field of requiredFields) {
      if (!paymentData[field]) {
        throw new Error(`Missing required field: ${field}`);
      }
    }

    return paymentData;
  }
}

module.exports = PaymentParser;