create table contas (
  conta_numero int,
  agencia_numero int,
  titular varchar(100), 
  tipo character,
  identificador varchar(20),
  ativa boolean,
  primary key(agencia_numero,conta_numero)
);

create table cartoes (
  cartao_numero numeric(18),
  agencia_numero int,
  conta_numero int,
  mes_ano_venc varchar(7),
  cvc int,
  limite decimal(12,2), 
  limite_disponivel decimal(12,2),
  ativo boolean,
  bloqueado boolean,
  FOREIGN KEY (agencia_numero,conta_numero) REFERENCES contas (agencia_numero,conta_numero),
  primary key(agencia_numero,conta_numero,cartao_numero)
);

insert into contas(conta_numero, agencia_numero, titular, tipo, identificador, ativa)
values
    (12345, 1234, 'Rodrigo Barbosa', 'F', '067.757.446-09', true),
    (23456, 2345, 'Thais Helena', 'F', '084.155.596-94', true),
    (34567, 3456, 'BnnCode', 'J', '60.813.719/0001-73', false);

insert into cartoes(cartao_numero, conta_numero, agencia_numero, cvc, mes_ano_venc, limite, limite_disponivel, ativo, bloqueado)
values
    (5134658357137960, 12345, 1234, 597, '01/30', 1000.00, 500.00, true, false),
    (5243779591303081, 23456, 2345, 011, '02/30', 2000.00, 1500.00, true, false),
    (4024007149542812, 12345, 1234, 666, '03/30', 500.00, 10.00, false, false),
    (344752177563906, 34567, 3456, 01, '04/30', 12000.00, 2500.00, false, false);

	