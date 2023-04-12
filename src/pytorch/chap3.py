import torch

t = torch.tensor([[[1,2.0],[3,4]],[[5,6],[7,8]]])

print(t.size())

print(t[0])

print(t.shape)

# 指定维度的张量平均数
print(t.mean(0))
print(t.mean(1))
print(t.mean(2))

t = torch.rand(2,3,4)
# 增加一个维度
t1 = t.unsqueeze(1)

# 在可以的情况下，压缩一个维度
print(t1.size())
t2 = t1.squeeze(0)
print(t2.size())
t3 = t1.squeeze(1)
print(t3.size())
t4 = t1.squeeze(2)
print(t4.size())
v_named = torch.rand(2,3,4,5,6,7).refine_names(..., "x", "y", "z", "a", "b", "c")
t_named = t.refine_names(..., "x", "y", "z")
# align_as返回一个添加了缺失维度的张量
t_align = t_named.align_as(v_named)
print("align", t_align.shape, t_align.names)

vt_multiply = v_named * t_align

print(vt_multiply.size())

vt_sum_multiply = vt_multiply.sum("y")
print(vt_sum_multiply.size(), vt_sum_multiply.names)

storage = vt_sum_multiply.untyped_storage()
print(storage)