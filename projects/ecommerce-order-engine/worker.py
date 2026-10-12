"""
Celery Background Worker for E-Commerce Engine
Handles async payment charging, invoice generation, inventory sync, and email dispatch.
"""

import time
import logging

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger("celery-worker")

def process_order_payment_and_fulfillment(order_id: str, user_id: str, amount: float):
    logger.info(f"==> Processing payment for order {order_id} (amount: ${amount}) via Stripe API")
    time.sleep(0.1) # Simulate external SaaS call
    logger.info(f"==> Payment captured successfully for order {order_id}")
    logger.info(f"==> Dispatching order receipt email to {user_id}")
    return {"status": "completed", "order_id": order_id, "paid": True}

if __name__ == "__main__":
    logger.info("Celery worker starting in test simulation mode...")
    res = process_order_payment_and_fulfillment("ord_test_100", "usr_test", 149.99)
    print("Worker result:", res)
