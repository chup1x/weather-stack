from loguru import logger
import request as req

welcome_msg = (
        "Доступные команды:\n"
        "/register - регистрация\n"
        "/weather - прогноз погоды\n"
        "/news - последние новости\n"
        "/profile - ваш профиль\n"
        "/clothes - рекомендации по одежде\n"
        "/help - помощь\n"
        "Вы можете получать информацию без регистрации\n"
        "Но каждый раз вводить данные заново\n"
    )

async def register(message, data) -> bool:
    user = message.from_user
    user_id = user.id
    username = user.username or user.first_name
    
    logger.info(f"регистрация {username} (ID: {user_id})")
    
    welcome_msg_error = '❌ Ошибка при регистрации'
    welcome_msg_local = (
        f"🎉 Добро пожаловать, {username}!\n"
        "Вы успешно зарегистрированы в системе!\n\n"
    )
    welcome_msg_local+= welcome_msg

    if user_exists(user_id): 
        logger.warning(f"{user_id} уже зарегистрирован")
        return False

    try:
        data[0] = username
        logger.debug(f"отправка данных регистрации для {user_id}: {data}")
        r = req.post_register_user(user_id, data[:])
        
        if r.status_code != 200:
            logger.error(f"ошибка регистрации: статус {r.status_code}")
            raise AssertionError(f"HTTP {r.status_code}")
        else:
            logger.debug("регистрация успешна, запрос профиля")
            r2 = req.get_user_profile(user_id)
            try:
                assert r2[0] == 200
                logger.success(f"{username} успешно зарегистрирован")
                await message.reply_text(welcome_msg_local)
                return True
            except AssertionError as e:
                logger.error(f"ошибка получения профиля после регистрации: {e}")
                await message.reply_text(welcome_msg_error)
                return False
            
    except AssertionError as e:
        logger.warning(f"Ошибка при регистрации: {e}")
        raise AssertionError("не удалось зарегистрировать пользователя")
    except Exception as e:
        logger.error(f"Неизвестная ошибка при регистрации: {e}")
        await message.reply_text(welcome_msg_error)
        return False

def user_exists(user_id: str) -> bool:
    logger.debug(f"проверка существования {user_id}")
    
    try:
        r2 = req.get_user_profile(user_id)
        if r2[0] == 200:
            logger.debug(f"{user_id} существует")
            return True
        else:
            logger.debug(f"{user_id} не найден, статус: {r2[0]}")
            raise AssertionError(f"HTTP {r2[0]}")
    except AssertionError as e:
        logger.warning(f"ошибка при проверке пользователя: {e}")
    except Exception as e:
        logger.error(f"Неизвестная ошибка при проверке пользователя: {e}")
        return False

def get_user_profile(user_id):
    logger.debug(f"запрос профиля {user_id}")
    
    try:
        r2 = req.get_user_profile(user_id)
        if r2[0] == 200:
            logger.debug(f"профиль {user_id} получен из апи")
            return r2[1]
        else:
            logger.warning(f"ошибка при получении профиля: {r2[0]}")
            raise AssertionError(f"HTTP {r2[0]}")
    except AssertionError as e:
        logger.warning(f"ошибка апи, используем локальную БД: {e}")
    except Exception as e:
        logger.error(f"Непонятная ошибка при получении профиля: {e}")
        return None