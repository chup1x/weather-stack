from loguru import logger
import os
import requests_async as requests
from dotenv import load_dotenv
load_dotenv()
url = os.getenv("BASE_URL")

async def post_register_user(usr_id: str, data):
    
    address = "profile/register"
    payload = {
        "name": str(data[0]),
        "sex": "male",
        "age": 69,
        "city_n": str(data[1]),
        "city_w": str(data[5]),
        "drop_time": str(data[6]),
        "t_comfort": int(data[2]),
        "t_tol": int(data[3]),
        "t_puh": int(data[4]),
        "telegram_id": int(usr_id)
    }

    try:
        logger.debug(f"отправка POST запроса на {url + address}")
        response = await requests.post(url + address, json=payload, timeout=10)
        logger.info(f"ответ сервера при регистрации: {response.status_code}")
        return response
    except Exception as e:
        logger.error(f"ошибка отправки данных регистрации: {e}")
        raise

async def get_user_profile(usr_id: str):
    
    address = f"profile/by-telegram-id/{usr_id}"

    try:
        logger.debug(f"отправка GET запроса на {url + address}")
        response = await requests.get(url + address, timeout=10)
        logger.info(f"ответ сервера при запросе профиля: {response.status_code}")
        
        user_data = None
        if response.status_code == 200:
            try:
                user_data = response.json()
                logger.debug(f"данные профиля получены: {user_data}")
            except Exception as e:
                logger.error(f"ошибка парсинга (профиль): {e}")
        else:
            logger.warning(f"неуспешный статус при запросе профиля: {response.status_code}")

        return response.status_code, user_data
        
    except Exception as e:
        logger.error(f"ошибка при запросе профиля пользователя: {e}")
        return None, None

async def get_weather(city: str):
    
    address = f"weather/city"

    try:
        response = await requests.get(url + address, data={'city':city}, timeout=10)
        logger.info(f"ответ сервера при запросе погоды по городу: {response.status_code}")
        
        user_data = None
        if response.status_code == 200:
            try:
                user_data = response.json()
                logger.debug("данные погоды по городу получены")
            except Exception as e:
                logger.error(f"ошибка парсинга (погода): {e},{response.text}")
        else:
            logger.warning(f"неуспешный статус при запросе погоды по городу: {response.status_code}")

        return response.status_code, user_data
        
    except Exception as e:
        logger.error(f"ошибка при запросе погоды по городу: {e}")
        return None, None

async def get_weather_with_profile(user_id: str):
    
    address = f"weather/by-telegram-id/{user_id}"

    try:
        response = await requests.get(url + address, timeout=10)
        logger.info(f"ответ сервера при запросе погоды через профиль: {response.status_code}")
        
        user_data = None
        if response.status_code == 200:
            try:
                user_data = response.json()
                logger.debug("данные погоды через профиль получены")
            except Exception as e:
                logger.error(f"ошибка парсинга (погода/профиль): {e},{response.text}")
        else:
            logger.warning(f"неуспешный статус при запросе погоды через профиль: {response.status_code}")

        return response.status_code, user_data
        
    except Exception as e:
        logger.error(f"ошибка при запросе погоды через профиль: {e}")
        return None, None

async def get_clothes_with_profile(user_id: str):
    
    address = f"weather/clothes/{user_id}"

    try:
        response = await requests.get(url + address, timeout=60)
        logger.info(f"ответ сервера при запросе одежды через профиль: {response.status_code}")
        
        user_data = None
        if response.status_code == 200:
            try:
                user_data = response.json()
                logger.debug("данные одежды через профиль получены")
            except Exception as e:
                logger.error(f"Ошибка парсинга (одежда/профиль): {e}")
        else:
            logger.warning(f"неуспешный статус при запросе одежды через профиль: {response.status_code}")

        return response.status_code, user_data
        
    except Exception as e:
        logger.error(f"ошибка при запросе одежды через профиль: {e}")
        return None, None


async def get_news(city: str):
    
    address = f"news/city"

    try:
        response = await requests.get(url + address, data={'city':city}, timeout=10)
        logger.info(f"ответ сервера при запросе новостей: {response.status_code}")
        
        user_data = None
        if response.status_code == 200:
            try:
                user_data = response.json()
                logger.debug("данные новостей получены")
            except Exception as e:
                logger.error(f"ошибка парсинга (новости): {e}")
        else:
            logger.warning(f"неуспешный статус при запросе новостей: {response.status_code}")

        return response.status_code, user_data
        
    except Exception as e:
        logger.error(f"ошибка при запросе новостей: {e}")
        return None, None

async def get_news_by_profile(user_id: str):
    
    address = f"news/by-telegram-id/{user_id}"

    try:
        response = await requests.get(url + address, timeout=10)
        logger.info(f"ответ сервера при запросе новостей по профилю: {response.status_code}")
        
        user_data = None
        if response.status_code == 200:
            try:
                user_data = response.json()
                logger.debug("данные новостей получены")
            except Exception as e:
                logger.error(f"ошибка парсинга (новости): {e}")
        else:
            logger.warning(f"неуспешный статус при запросе новостей: {response.status_code}")

        return response.status_code, user_data
        
    except Exception as e:
        logger.error(f"ошибка при запросе новостей: {e}")
        return None, None
