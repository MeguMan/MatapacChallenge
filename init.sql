create table users (
    tg_id          bigint  unique primary key,
    tg_username    text    not null,
    sol_public_key text    unique not null,
    attempt        int     default 1,
    created_at  timestamp without time zone default NOW()
);

insert into users
    (tg_id, tg_username, sol_public_key)
VALUES
    (5029548180,'umeelazy','A3QqMrnvFvDeiscbUa8C6gmh67wCbCfjZpSS3e1ZNq2y'),
    (643393737,'Provod1415','2XpciFhbX4cZbB4fp4WMJVdqySeShgeqe44So8BRkFve'),
    (1118706387,'temakz1','2Lz158H98eqh3FR69qKCv3PPwmzrvipQw3Ep5sUazNnc'),
    (6846018026,'freedom_is_comin','HgVvb24SR6bzXjpEshiYMf1QtQpw8dohQCyMsMeHhiqX'),
    (6126403593,'vivarium6137','CCoC7P4qotmMfb13B8DUSUAZkA8vVVzWKADRqcMqjrjQ'),
    (5562975425,'ttimcr','5SeUJX2p3AzsB18yyfz4565a2gzVM5yFBhnQCtfUUUgg'),
    (1580491880,'marioner1','4BucDEGTyzQNvxjhxqqb9WWE1pX71SdHTq2ZNGhCsntu'),
    (714294900,'chouqxx','HkXSodmSYVFyAuVjRYsmfmZjQbr6W6aksHiUVYsKhagN'),
    (75002370,'RMNHTRMN','4BWadTHkrR1Vto6Bv8BAEyhaf8MAoQU4uKQRHfjcvmKa'),
    (748763882,'Andryusha2','CCHLt6SCyLPxotYbLju18M5uN6fHazea35ZiiU9JEkwD'),
    (445289901,'iFedotspb','9MakS2Uq9Qrp1WsRbbBg1XQry2qS1XpYcnm1JMBJL94p'),
    (5488544766,'odindva1','4KN6t5UnfVFoFFLpoUb2ToVsqusLi4SsYbuuFdKvmKtK'),
    (1172467181,'Ziuom','DV1RpBRMoywjqe3bpYLmAd3ve8yf2CzztinVAGXeUhgJ'),
    (820194641,'AntonSushkov','XXXYxsvw1NG2s3S2UZ9v3TDBABo6wsCB512DzL37mr1'),
    (893930741,'Vovchikk_k','ECZKniJcdtDcbAVodgq4JicugVujg5C3EPJegyLh8s6u'),
    (247355108,'bateshOn','8kmGD3jUvSpy3CC49JTN4e1hZ8kJf2d43oxagjuCxDCw'),
    (5831235772,'ijweifjw','6ZM3RZLTh7YXeTRiKJ8Gw4nsKDJUuF6BRw8MVGKvCZZg'),
    (5511930657,'r3gul4rm4n','FMmvgji4ycuF8tbuisaKzt7NZu4uFaavJVkp3cKecwRk'),
    (750007301,'abzdng','FC99Bg3WnwUQC3QZECih5SKzUQfSaHjTZkkyUycjetUn'),
    (5674065860,'us3ro_0','FAtg37XoRt3RyPc5Tx9JiypnCjKmnX2C4hh7E2VcAtvE'),
    (500077713,'Injuste','pBG83nep2eJBKEEMdZ1eJY7pqT3mhESfL5gV3SZ7FQj'),
    (1641618614,'eg0rkakryg','Gu5Bys9gHbsoqKS7xKQoCfDSBQ23De9m6SJFTth7Qktk'),
    (1709763863,'ibraxim_m','Bm83SzDEfewP8NWmW2x8tRH2Z4BizUdZcQPFwqfjGdZ2'),
    (1364701840,'Bonifatia','9MXyr7Z82aU5Fw6H6z9DF6J1bKvKpvdtEXD5mPuLxdrG'),
    (5975937856,'','2xonuQMAHzwPM7LJaMnpSBHhPQgLpf52CqjGwp7shMeV'),
    (351469103,'yououma','EPECBC5wrxRknsrxmDnBJ3fv9LR3wAnvx6kYPa92nU5r'),
    (213506548,'KARM404','o8ot7WmSZa1VU5EKxpHn62XpPGenLaFypS3FV5N4xJG'),
    (1324194431,'morsxdd','CwGoM2Fff6V5Czs8QaruCfewS2ChGH5ubpUoQGYsDswF');