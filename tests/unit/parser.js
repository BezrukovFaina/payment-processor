const fs = require('fs');
const path = require('path');

class PaymentParser {
  constructor(filePath) {
    this.filePath = filePath;
    this.payments = [];
  }

  parse() {
    try {
      const data = fs.readFileSync(this.filePath, 'utf8');
      const lines = data.split('\n');
      lines.forEach((line) => {
        const payment = this.parsePayment(line);
        if (payment) {
          this.payments.push(payment);
        }
      });
    } catch (error) {
      throw new Error(`Failed to parse payments: ${error.message}`);
    }
  }

  parsePayment(line) {
    const regex = /^(\d{4}-\d{2}-\d{2})\s+(\d{1,10})\s+(\w+)\s*$/;
    const match = line.match(regex);
    if (match) {
      return {
        date: match[1],
        amount: parseFloat(match[2]),
        type: match[3]
      };
    }
    return null;
  }

  getPayments() {
    return this.payments;
  }
}

module.exports = PaymentParser;