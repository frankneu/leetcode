import torch
import torch.nn as nn
import torch.optim as optim

# 裸写一个线性模型：定义数据，定义模型，定义train
model = nn.Linear(20, 1)
optimizer = optim.SGD(model.parameters(), lr = 1e-2)

def train(epochs, model, loss, optimizer, train_param, train_value, valid_param, valid_value):
    for epoch in range(1, epochs + 1):
        train_predict = model(train_param)
        train_loss = loss(train_predict, train_value)

        optimizer.zero_grad()
        optimizer.step()
        train_loss.backward()

        if 0 == epochs % 10:
            valid_predict = model(valid_param)
            valid_loss = loss(valid_predict, valid_value)
            print(f"Epoch {epoch}, Training loss {train_loss.item():.4f},"
                  f" Validation loss {valid_loss.item():.4f}")

