# 原始的全数据表
f = open('/Users/frank/Documents/202306/ruanduo_migrate/think_apply.sql', 'r')
source_sql = f.readlines()
f.close()

# 支付宝小程序云的库列表
f = open('/Users/frank/Downloads/mendian-store', 'r')
cloud_db_mdst = f.readlines()
f.close()

f = open('/Users/frank/Downloads/mendian-shop', 'r')
cloud_db_mdsp = f.readlines()
f.close()

f = open('/Users/frank/Downloads/danbangtuan-top', 'r')
cloud_db_dbt = f.readlines()
f.close()

f = open('/Users/frank/Downloads/cloudqxt-com', 'r')
cloud_db_cq = f.readlines()
f.close()

mdst_source_database_to_site_from_sql = dict()
mdsp_source_database_to_site_from_sql = dict()
dbt_source_database_to_site_from_sql = dict()
cq_source_database_to_site_from_sql = dict()
for i in source_sql:
    i = i.replace("'", "")
    array = i.split(",")
    if len(array) > 9:
        if array[1].strip().endswith("mendian.store"):
            mdst_source_database_to_site_from_sql[array[9].strip()] = array[1].strip()
            continue
        if array[1].strip().endswith("mendian.shop"):
            mdsp_source_database_to_site_from_sql[array[9].strip()] = array[1].strip()
            continue
        if array[1].strip().endswith("danbangtuan.top"):
            dbt_source_database_to_site_from_sql[array[9].strip()] = array[1].strip()
            continue
        if array[1].strip().endswith("cloudqxt.com"):
            cq_source_database_to_site_from_sql[array[9].strip()] = array[1].strip()
            continue

mdst_cloud_database = set()
for i in cloud_db_mdst:
    mdst_cloud_database.add(i.strip())

mdsp_cloud_database = set()
for i in cloud_db_mdsp:
    mdsp_cloud_database.add(i.strip())

dbt_cloud_database = set()
for i in cloud_db_dbt:
    dbt_cloud_database.add(i.strip())

cq_cloud_database = set()
for i in cloud_db_cq:
    cq_cloud_database.add(i.strip())

print("mendian.store~~~~~~")
miss_mdst_cloud_database = set()
for i in mdst_source_database_to_site_from_sql.keys():
    if not mdst_cloud_database.__contains__(i):
        miss_mdst_cloud_database.add(i)
        print(i)
print("','".join(miss_mdst_cloud_database))
print("mendian.shop~~~~~~")
miss_mdsp_cloud_database = set()
for i in mdsp_source_database_to_site_from_sql.keys():
    if not mdsp_cloud_database.__contains__(i):
        miss_mdsp_cloud_database.add(i)
        print(i)
print("','".join(miss_mdsp_cloud_database))
print("danbangtuan.top~~~~~~")
miss_dbt_cloud_database = set()
for i in dbt_source_database_to_site_from_sql.keys():
    if not dbt_cloud_database.__contains__(i):
        miss_dbt_cloud_database.add(i)
        print(i)
print("','".join(miss_dbt_cloud_database))
print("cloudqxt.com~~~~~~")
miss_cq_cloud_database = set()
for i in cq_source_database_to_site_from_sql.keys():
    if not cq_cloud_database.__contains__(i):
        miss_cq_cloud_database.add(i)
        print(i)
print("','".join(miss_cq_cloud_database))