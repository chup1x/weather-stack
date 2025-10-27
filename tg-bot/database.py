import sqlite3 as sq

alf = ['name', 'sex', 'age', 'temp1', 'temp2', 'temp3', 'news', 'right']

class user_pack():
    
    def __init__(self, columns: list[int],*values):
        m = [0]*8
        for i,v in enumerate(columns): m[v] = values[i]
    
        self.d = {
            'name' : m[0], 'sex' : m[1] , 'age' : m[2],
            'temp1': m[3], 'temp2' : m[4], 'temp3' : m[5],
            'news' : m[6], 'right' : m[7]
            }
    def inone(self,col: str, value): self.d[col]=value
    def inmany(self,columns: list[int],*values):
        keys = list(self.d.keys())
        for i,v in enumerate(columns): self.d[keys[v]] = values[i]
    def unpack(self): return '(' + str(list(self.d.values()))[1:-1] + ')'
    def get_pack(self): return list(self.d.values())
        
class session():
    
    def __init__(self): 
        self.con = sq.connect('NWBot.db', check_same_thread=False)
        self.cur = self.con.cursor()
        self.create_table()

    def create_table(self):
        """Создает таблицу если она не существует"""
        self.cur.execute('''
            CREATE TABLE IF NOT EXISTS user_params (
                id INTEGER PRIMARY KEY,
                name TEXT,
                sex TEXT,
                age INTEGER,
                temp1 REAL,
                temp2 REAL,
                temp3 REAL,
                news TEXT,
                right TEXT
            )
        ''')
        self.con.commit()

    #####################################################
    #   ОТРИСОВЫВЕТ ВСЕ ЗАПИСИ В ТАБЛИЦЕ                #
    #   ПРИНИМАЕТ СТРОКУ СОРТИРОВКИ (МОЖНО '')          #
    #   И ИМЕНА КОЛОНОК ЕСЛИ НАДО                       #
    #####################################################
    def print_(self, sort: str, *num):
        if len(num)>0:
            sym = ",".join(num)
        else: sym = '*'
        data = 'SELECT ' + sym + ' FROM user_params' + sort
        print(self.cur.execute(data).fetchall())


    #####################################################
    #   ВЫДАЁТ ОДНУ ПОЗИЦИЮ ПО ID ИЛИ ИМЕНИ             #
    #   ПРИНИМАЕТ НОМЕРА КОЛОНОК (ОБЯЗАТЕЛЬНО) И ID/NAME#
    #####################################################
    def getby_id(self, columns: list[int], idd: int) -> list:
        clms = ','.join([alf[i] for i in columns])
        data = 'SELECT ' + clms + f' FROM user_params WHERE id = {idd}'
        return self.cur.execute(data).fetchone()
    
    def getby_name(self, columns: list[int], name: str) -> list:
        clms = ','.join([alf[i] for i in columns])
        data = 'SELECT ' + clms + f' FROM user_params WHERE name = "{name}"'
        return self.cur.execute(data).fetchone()


    #####################################################
    #   ВСТАВЛЯЕТ НОВУЮ ЗАПИСЬ ДВУМЯ СПОСОБАМИ:         #
    #   ЧЕРЕЗ УНИФИЦИРОВАННЫЙ USER-PACK                 #
    #   ВВОДОМ LIST КОЛОНОК ЗАТЕМ ПЕРЕЧИСЛЕНИЯ ДАННЫХ(,)#
    #####################################################
    def insrt_pack(self, pack: user_pack):
        clms = '(' + ','.join(alf)+')'
        value = pack.unpack()
        data = f'INSERT INTO user_params {clms} VALUES {value}'
        self.cur.execute(data)
        self.con.commit()
        
    def insrt_with_id(self, user_id: int, columns: list[int], *values):
        """Вставляет запись с указанным ID"""
        clms = 'id,' + ','.join([alf[i] for i in columns])
        value_str = str((user_id,) + values)
        data = f'INSERT INTO user_params ({clms}) VALUES {value_str}'
        self.cur.execute(data)
        self.con.commit()
        
    def insrt_(self, columns: list[int],*values):
        clms = '(' + ','.join([alf[i] for i in columns]) +')'
        value = str(values)
        data = f'INSERT INTO user_params {clms} VALUES {value}'
        self.cur.execute(data)
        self.con.commit()

    #####################################################
    #   РАБОТАЕТ АНАЛОГИЧНО ОДИНАРНОМУ INSERT           #
    #   ПРИНИМАЕТ ЛИБО LIST ЛИБО USER-PACK              #
    #   ПРИМЕР МАССИВА БЕЗ ИСПОЛЬЗОВАНИЯ USER-PACK:     #
    #   FAN = [("Tem",1,2),("Yan",3,4),("Dan",5,6)]     #
    #####################################################
    def insrt_many_columns(self, columns: list[int], dat):
        q = ('?, '*len(columns))[:-2]
        clms = '(' + ','.join([alf[i] for i in columns]) +')'
        data = f'INSERT INTO user_params {clms} VALUES({q})'
        self.cur.executemany(data, dat)
        self.con.commit()
        
    def insrt_many_packs(self, packs: list[user_pack]):
        clms = '(' + ','.join(alf)+')'
        values = [pack.get_pack() for pack in packs]
        q = ('?, '*len(alf))[:-2]
        data = f'INSERT INTO user_params {clms} VALUES({q})'
        self.cur.executemany(data, values)
        self.con.commit()

    #####################################################
    #   ИЗМЕНЯЕТ СОСТОЯНИЕ ВЫБРАННЫХ ЯЧЕЕК В ЗАПИСИ     #
    #   ПРИИНИМАЕТ LIST КОЛОНОК И ЕЩЁ ЧЕТО, ПОКА НЕ РОБИ#
    #####################################################
    def updatecl(self, columns: list[int], user_id: int, values: list):
        set_clause = ','.join([f'{alf[col]} = ?' for col in columns])
        data = f'UPDATE user_params SET {set_clause} WHERE id = {user_id}'
        self.cur.execute(data, values)
        self.con.commit()

    #####################################################
    #   УДАЛЯЕТ ЗАПИСИ В ТАБЛИЦЕ                        #
    #   ПРИНИМАЕТ ПАРАМЕТР ПОИСКА ЗАПИСИ В ВИДЕ STR TEXT#
    #####################################################
    def delr(self, param: str):
        data = f'DELETE FROM user_params WHERE {param}'
        self.cur.execute(data)
        self.con.commit()

    def user_exists(self, user_id: int) -> bool:
        """Проверяет существует ли пользователь с указанным ID"""
        data = f'SELECT id FROM user_params WHERE id = {user_id}'
        result = self.cur.execute(data).fetchone()
        return result is not None

    def get_user_profile(self, user_id: int):
        """Получает полный профиль пользователя"""
        data = f'SELECT * FROM user_params WHERE id = {user_id}'
        result = self.cur.execute(data).fetchone()
        return result

    def close(self):
        """Закрывает соединение с БД"""
        self.con.close()
