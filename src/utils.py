# utils.py

import logging
import uuid
from datetime import datetime
from payment_processor.config import Config

class Utils:
    @staticmethod
    def generate_payment_id(length=10):
        return str(uuid.uuid4().int)[:length]

    @staticmethod
    def get_current_date():
        return datetime.now().strftime('%Y-%m-%d %H:%M:%S')

    @staticmethod
    def email_template(name, amount, payment_id):
        return f'Dear {name},\n\nYour payment of ${amount} has been processed.\nPayment ID: {payment_id}\n\nThank you for using our service.\n'

    @staticmethod
    def log_message(level, message):
        logging.basicConfig(filename=Config.LOG_FILE, level=Config.LOG_LEVEL)
        logging.log(level, message)

    @staticmethod
    def validate_amount(amount):
        try:
            amount = float(amount)
            if amount <= 0:
                raise ValueError
            return amount
        except ValueError:
            raise ValueError('Invalid amount')