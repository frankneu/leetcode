from matplotlib import pyplot as plt
import torch
def model(tu, w, b):
    return w * tu + b

def loss_fn(tp, tc):
    return tp-tc

param = torch.tensor([1.0, 0.0], requires_grad=True)
loss = loss_fn(model())