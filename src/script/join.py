# 原始的全数据表
f = open('/Users/frank/Downloads/sql_result', 'r')
source_sql = f.readlines()
f.close()

cloud_database = set()
for i in source_sql:
    if not i.startswith("Table"):
        cloud_database.add(i.strip())

for i in cloud_database:
    print(i)


