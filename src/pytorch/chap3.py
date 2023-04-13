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

points = vt_sum_multiply[1]

print("before", points.size())
# stride是在指定维度dim中从一个元素跳到下一个元素所必需的步长。当没有参数传入时，返回所有步长的元组。否则，将返回一个整数值作为特定维度dim中的步长。
# （5*6*7, 6*7, 7, 1）
stride = points.stride()

print("stride", stride)

test_trans = torch.rand(2,3,4,5)
print(test_trans.size())
transposed = test_trans.transpose(0,2)
print(transposed.size())
# 某些Tensor操作（如transpose、permute、narrow、expand）与原Tensor是共享内存中的数据，不会改变底层数组的存储，但原来在语义上相邻、内存里也相邻的元素在执行这样的操作后，在语义上相邻，但在内存不相邻，即不连续了（is not contiguous）。
print(test_trans.is_contiguous())
print(transposed.is_contiguous())
contiguous = transposed.contiguous()
print(contiguous.is_contiguous())

torch.save(contiguous, "contiguous.t")

import h5py
file_w = h5py.File("contiguous.hdf5", "w")
file_w.create_dataset("contiguous", data = contiguous.numpy())
file_w.close()

file_r = h5py.File("contiguous.hdf5", "r")
file_read_numpy = file_r["contiguous"]
# ???
tensor_from_h5py = torch.from_numpy(file_read_numpy[-2 :])
print(tensor_from_h5py.size())
file_r.close()