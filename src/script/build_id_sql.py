f = open('/Users/frank/Documents/ruanduo_mendian_shop_sql_id/database_list', 'r')
database_list = f.readlines()
f.close()

f = open('/Users/frank/Documents/ruanduo_mendian_shop_sql_id/table_list', 'r')
table_list = f.readlines()
f.close()

f = open('/Users/frank/Downloads/mendian-shop-id-list-202307212110', 'r')
id_result = f.readlines()
f.close()


shop_database_list = list()
for i in database_list:
    shop_database_list.append(i.strip())

shop_table_set = set()
for i in table_list:
    shop_table_set.add(i.strip())

id_result_list = list()
for i in id_result:
    id_result_list.append(i)


response = dict()
database = ""
table_id_map = dict()
id_index = 0
while id_index < len(id_result_list):
    value = id_result_list[id_index]
    array = value.split("______")
    if value.startswith("database()"):
        id_index += 1
        database = id_result_list[id_index]
        response[database.strip()] = dict()
        table_id_map = response[database.strip()]
        id_index += 1
        continue
    if array.__len__() == 2 and shop_table_set.__contains__(array[1].strip()):
        id_index += 1
        id = id_result_list[id_index]
        table_id_map[value.strip()] = id
        id_index += 1
        continue
    id_index += 1

f = open('mendian-shop-mysqldump.sh', 'w')
for database in response.keys():
    table_id_map = response[database]
    for key in table_id_map.keys():
        array = key.split("______")
        column = array[0]
        table = array[1]
        f.write("mysqldump -uadmin -hmysql-mendian-shop -p "
                + database + " " + table
                + "--where \"" + column + ">" + table_id_map[key])

f.close()




