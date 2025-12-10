from logger_config import logger
from telegram.ext import Application
from handlers import get_handlers

def main() -> None:
    logger.info("Запуск")
    
    try:
        TELEGRAM_TOKEN = "7958428332:AAE0tAMv44v4pmMB-mzydqRN8ppBv6FIlrs"
        if not TELEGRAM_TOKEN:
            logger.error("TELEGRAM_TOKEN не найден")
            raise ValueError("TELEGRAM_TOKEN не найден")
        
        application = Application.builder().token(TELEGRAM_TOKEN).build()

        handlers = get_handlers()
        for handler in handlers:
            application.add_handler(handler)

        logger.success(f"Успешно загружено обработчиков: {len(handlers)}")
        logger.info("запущено")
        
        application.run_polling(
            drop_pending_updates=True,
            allowed_updates=["message", "callback_query"]
        )
        
    except Exception as error:
        logger.critical(f"ошибка при запуске бота: {error}")
        raise

if __name__ == '__main__':
    main()
