import telebot
import requests
from datetime import datetime, timedelta
from database import session, user_pack

# Инициализация бота и БД
bot = telebot.TeleBot('7958428332:AAE0tAMv44v4pmMB-mzydqRN8ppBv6FIlrs')
db = session()

user_states = {}
NEWS_API_KEY = '0fac40f7dcd34967af176019e1c6a526'

@bot.message_handler(commands=['start'])
def main(message):
    user_id = message.from_user.id
    username = message.from_user.username or message.from_user.first_name
    
    print(f"Проверяем пользователя {user_id} ({username})")  # Для отладки
    
    # Проверяем, зарегистрирован ли пользователь
    if not db.user_exists(user_id):
        print(f"Пользователь {user_id} не найден, регистрируем...")  # Для отладки
        
        # Регистрируем нового пользователя с его Telegram ID
        try:
            # Используем метод с явным указанием ID
            db.insrt_with_id(user_id, [0], username)
            print(f"Пользователь {user_id} успешно зарегистрирован")  # Для отладки
            
            welcome_msg = f'''🎉 Добро пожаловать, {username}!
Вы успешно зарегистрированы в системе!

Доступные команды:
/weather - прогноз погоды
/news - последние новости
/profile - ваш профиль
/help - помощь'''
        except Exception as e:
            print(f"Ошибка при регистрации: {e}")  # Для отладки
            welcome_msg = f'❌ Ошибка при регистрации: {e}'
    else:
        print(f"Пользователь {user_id} уже существует")  # Для отладки
        welcome_msg = f'''👋 С возвращением, {username}!

Чем могу помочь?
/weather - прогноз погоды
/news - последние новости
/profile - ваш профиль
/help - помощь'''
    
    bot.send_message(message.chat.id, welcome_msg, parse_mode='html')

@bot.message_handler(commands=['help'])
def help_command(message):
    help_text = '''
📋 Доступные команды:

/start - регистрация и начало работы
/weather - прогноз погоды для вашего города
/news - последние новости по интересующему городу
/profile - информация о вашем профиле

💡 Просто введите команду и следуйте инструкциям!
'''
    bot.send_message(message.chat.id, help_text, parse_mode='html')

@bot.message_handler(commands=['profile'])
def profile_command(message):
    user_id = message.from_user.id
    
    print(f"Запрос профиля для пользователя {user_id}")  # Для отладки
    
    if not db.user_exists(user_id):
        bot.send_message(message.chat.id, 
                        '❌ Вы не зарегистрированы. Используйте /start для регистрации.')
        return
    
    profile = db.get_user_profile(user_id)
    
    if profile:
        profile_text = f'''
📊 Ваш профиль:

🆔 ID: {profile[0]}
👤 Имя: {profile[1] or 'Не указано'}
🚻 Пол: {profile[2] or 'Не указан'}
🎂 Возраст: {profile[3] or 'Не указан'}
🌡️ Температура 1: {profile[4] or 'Нет данных'}
🌡️ Температура 2: {profile[5] or 'Нет данных'}
🌡️ Температура 3: {profile[6] or 'Нет данных'}
📰 Последние новости: {profile[7] or 'Нет данных'}
🔑 Права: {profile[8] or 'user'}
'''
    else:
        profile_text = '❌ Профиль не найден.'
    
    bot.send_message(message.chat.id, profile_text, parse_mode='html')

@bot.message_handler(commands=['weather'])
def weather_command(message):
    user_id = message.from_user.id
    
    if not db.user_exists(user_id):
        bot.send_message(message.chat.id, 
                        '❌ Вы не зарегистрированы. Используйте /start для регистрации.')
        return
    
    bot.send_message(message.chat.id, 
                    '🌤️ Введите название вашего города для получения прогноза погоды:')
    user_states[message.chat.id] = 'waiting_city'

@bot.message_handler(commands=['news'])
def news_command(message):
    user_id = message.from_user.id
    
    if not db.user_exists(user_id):
        bot.send_message(message.chat.id, 
                        '❌ Вы не зарегистрированы. Используйте /start для регистрации.')
        return
    
    bot.send_message(message.chat.id, 
                    '📰 Введите интересующий вас город для поиска новостей:')
    user_states[message.chat.id] = 'waiting_city_news'

def get_weather(city, chat_id):
    try:
        url = f'https://api.openweathermap.org/data/2.5/weather?q={city}&units=metric&lang=ru&appid=79d1ca96933b0328e1c7e3e7a26cb347'
        response = requests.get(url)
        weather_data = response.json()
        
        if response.status_code == 200:
            temperature = weather_data['main']['temp']
            description = weather_data['weather'][0]['description']
            humidity = weather_data['main']['humidity']
            pressure = weather_data['main']['pressure']
            wind_speed = weather_data['wind']['speed']
            
            weather_message = (f"🌤️ Погода в городе {city}:\n"
                             f"🌡️ Температура: {temperature}°C\n"
                             f"📝 Описание: {description}\n"
                             f"💧 Влажность: {humidity}%\n"
                             f"📊 Давление: {pressure} гПа\n"
                             f"💨 Скорость ветра: {wind_speed} м/с")
            
            bot.send_message(chat_id, weather_message, parse_mode='html')
            
            # Сохраняем температуру в БД
            user_id = chat_id
            if db.user_exists(user_id):
                # Получаем текущие данные о температуре
                user_data = db.getby_id([3, 4, 5], user_id)
                if user_data:
                    temp1, temp2, temp3 = user_data
                    # Сдвигаем температуры и добавляем новую
                    new_temps = [temperature, temp1, temp2]
                    db.updatecl([3, 4, 5], user_id, new_temps)
            
        else:
            bot.send_message(chat_id, 
                           '❌ Город не найден. Проверьте правильность написания.', 
                           parse_mode='html')
            
    except Exception as e:
        print(f"Ошибка: {e}")
        bot.send_message(chat_id, 
                       '❌ Произошла ошибка. Попробуйте еще раз.', 
                       parse_mode='html')

def get_news(city, chat_id):
    try:
        week_ago = (datetime.now() - timedelta(days=7)).strftime('%Y-%m-%d')
        
        url = f'https://newsapi.org/v2/everything?q={city}&from={week_ago}&sortBy=publishedAt&language=ru&apiKey={NEWS_API_KEY}'
        
        response = requests.get(url)
        news_data = response.json()
        
        if response.status_code == 200 and news_data['status'] == 'ok':
            articles = news_data['articles'][:5]
            
            if not articles:
                bot.send_message(chat_id, 
                               f'❌ Новости по городу {city} не найдены за последнюю неделю.', 
                               parse_mode='html')
                return
            
            news_message = f"📰 Последние новости по городу {city}:\n\n"
            
            for i, article in enumerate(articles, 1):
                title = article['title']
                source = article['source']['name']
                published_at = datetime.strptime(article['publishedAt'], '%Y-%m-%dT%H:%M:%SZ').strftime('%d.%m.%Y %H:%M')
                url = article['url']
                
                news_message += f"{i}. **{title}**\n"
                news_message += f"   📋 Источник: {source}\n"
                news_message += f"   🕒 Дата: {published_at}\n"
                news_message += f"   🔗 [Читать полностью]({url})\n\n"
            
            bot.send_message(chat_id, news_message, parse_mode='Markdown')
            
            # Сохраняем последний запрос новостей в БД
            user_id = chat_id
            if db.user_exists(user_id):
                db.updatecl([6], user_id, [city])
            
        else:
            error_msg = news_data.get('message', 'Неизвестная ошибка')
            bot.send_message(chat_id, 
                           f'❌ Ошибка при получении новостей: {error_msg}', 
                           parse_mode='html')
            
    except Exception as e:
        print(f"Ошибка при получении новостей: {e}")
        bot.send_message(chat_id, 
                       '❌ Произошла ошибка при получении новостей. Попробуйте еще раз.', 
                       parse_mode='html')

@bot.message_handler(func=lambda message: True)
def handle_all_messages(message):
    if message.chat.id in user_states:
        city = message.text.strip()
        
        if user_states[message.chat.id] == 'waiting_city':
            del user_states[message.chat.id]
            get_weather(city, message.chat.id)
            
        elif user_states[message.chat.id] == 'waiting_city_news':
            del user_states[message.chat.id]
            get_news(city, message.chat.id)
            
    else:
        bot.send_message(message.chat.id, 
                       '❌ Я не понимаю эту команду. Используйте /help для просмотра доступных команд.')

if __name__ == '__main__':
    try:
        print("🤖 Бот запущен...")
        bot.polling(non_stop=True)
    except KeyboardInterrupt:
        print("🛑 Бот остановлен")
    finally:
        db.close()
